package checker

import (
	"errors"
	"fmt"

	"github.com/fishmanDK/miet_project/pkg/logger"
	"github.com/jmoiron/sqlx"
	"gopkg.in/gomail.v2"

	sq "github.com/Masterminds/squirrel"
)

const workersCount = 10

var checkerIsNotStarted = errors.New("checker not running")
var checkerStopped = errors.New("checker is already stopped")

type Message struct {
	CassetteID int
	Count      int
}

type CheckerFirstReserveUsers struct {
	isStarted bool
	ch        chan Message
	db        *sqlx.DB
	log       logger.Logger
}

func NewCheckerFirstReserveUsers(db *sqlx.DB, log logger.Logger) *CheckerFirstReserveUsers {
	return &CheckerFirstReserveUsers{
		ch: make(chan Message, 100),
		db: db,
		log: log,
	}
}

func (c *CheckerFirstReserveUsers) Start() {
	c.isStarted = true

	for i := 0; i < workersCount; i++ {
		go func() {
			for v := range c.ch {
				data, err := c.getUsers(v.CassetteID, v.Count)
				if err != nil {
					c.log.Info(err.Error())
					continue
				}
				c.log.Info(fmt.Sprintf("Ожидают отправку: %v", data))

				for _, d := range data {
					err = c.sendEmail(d)
					if err != nil {
						c.log.Info(err.Error())
						continue
					}
					c.log.Info(fmt.Sprintf("Оповещение отправленно на почту: %s", d.email))
				}
			}
		}()
	}
}

func (c *CheckerFirstReserveUsers) Push(message Message) error {
	if !c.isStarted {
		return checkerIsNotStarted
	}

	c.ch <- message
	return nil
}

func (c *CheckerFirstReserveUsers) Stop() error {
	if !c.isStarted {
		return checkerStopped
	}

	close(c.ch)
	return nil
}

type dataToNotify struct {
	email         string `db:"email"`
	name          string `db:"name"`
	ganre         string `db:"genre"`
	yearOfRelease int    `db:"year_of_release"`
}

func (c *CheckerFirstReserveUsers) getUsers(cassetteID, count int) ([]dataToNotify, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return []dataToNotify{}, err
	}
	defer tx.Rollback()

	query, args, err := sq.
	Select("reserve_pool.cassette_id", "reserve_pool.user_id", "email", "name", "genre", "year_of_release").
	From("reserve_pool").
		Join("users ON reserve_pool.user_id = users.id").
		Join("cassettes ON reserve_pool.cassette_id = cassettes.id").
		Where(sq.Eq{"reserve_pool.cassette_id": cassetteID}).
		OrderBy("reserve_pool.reservation_date ASC").
		Limit(uint64(count)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return []dataToNotify{}, err
	}

	rows, err := tx.Query(query, args...)
	if err != nil {
		return []dataToNotify{}, err
	}
	defer rows.Close()

	var data []dataToNotify
	var recordsToDelete []struct {
		CassetteID int
		UserID     int
	} 

	for rows.Next() {
		var d dataToNotify
		var cassetteID, userID int

		err := rows.Scan(&cassetteID, &userID, &d.email, &d.name, &d.ganre, &d.yearOfRelease)

		if err != nil {
			return []dataToNotify{}, err
		}
		data = append(data, d)
	}

	if err := rows.Err(); err != nil {
		return []dataToNotify{}, err
	}

	if len(recordsToDelete) > 0 {
		deleteBuilder := sq.Delete("reserve_pool").PlaceholderFormat(sq.Dollar)

		for _, record := range recordsToDelete {
			deleteBuilder = deleteBuilder.Where(sq.And{
				sq.Eq{"cassette_id": record.CassetteID},
				sq.Eq{"user_id": record.UserID},
			})
		}

		deleteQuery, deleteArgs, err := deleteBuilder.ToSql()
		if err != nil {
			return []dataToNotify{}, err
		}

		_, err = tx.Exec(deleteQuery, deleteArgs...)
		if err != nil {
			return []dataToNotify{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return []dataToNotify{}, err
	}

	return data, nil
}

func (c *CheckerFirstReserveUsers) sendEmail(data dataToNotify) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "denis.23weer@gmail.com")
	m.SetHeader("To", data.email)
	m.SetHeader("Subject", "Поступление кассет")
	m.SetBody("text/html", fmt.Sprintf("Кассета %s (жанр: %s, год релиза: %d) уже доступна для заказа!)", data.name, data.ganre, data.yearOfRelease))

	d := gomail.NewDialer("smtp.gmail.com", 587, "denis.23weer@gmail.com", "bqlo xmrm ebrv ccle")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}