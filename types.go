package main

import (
	"database/sql"
)

type ChatListItem struct {
	Id                                sql.NullInt64  `sql:"_id"`
	KeyRemoteJId                      sql.NullString `sql:"key_remote_jid"`
	MessageTableId                    sql.NullInt64  `sql:"message_table_id"`
	Subject                           sql.NullString `sql:"subject"`
	Creation                          sql.NullInt64  `sql:"creation"`
	LastReadMessageTableId            sql.NullInt64  `sql:"last_read_message_table_id"`
	LastReadReceiptSentMessageTableId sql.NullInt64  `sql:"last_read_receipt_sent_message_table_id"`
	Archived                          sql.NullInt64  `sql:"archived"`
	SortTimestamp                     sql.NullInt64  `sql:"sort_timestamp"`
}
