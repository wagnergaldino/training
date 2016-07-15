package riskzone

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/wagnergaldino/training/dne/util"
)

func Execute(db *sql.DB) {
	restricoes, err := db.Query("select id, inicial, final, restricao from restriction where processado is null order by id")
	util.CheckErr(err)
	defer restricoes.Close()

	for restricoes.Next() {

		var id int
		var inicial, final, restricao string
		err = restricoes.Scan(&id, &inicial, &final, &restricao)
		util.CheckErr(err)

		inicial = util.FillBefore(inicial, "0", 8)
		final = util.FillBefore(final, "0", 8)

		ceps, err := db.Query("select cep from address where cep between " + inicial + " and " + final + " order by cep")
		util.CheckErr(err)

		for ceps.Next() {

			var cep string
			err = ceps.Scan(&cep)
			util.CheckErr(err)
			salvaRestricao(db, restricao, cep)

		}

		ceps.Close()
		salvaStatus(db, id)

	}

}

func salvaRestricao(db *sql.DB, restricao, cep string) {
	stmt, err := db.Prepare("update address set restricao = ? where cep = ?")
	util.CheckErr(err)
	_, err = stmt.Exec(restricao, cep)
	util.CheckErr(err)
	stmt.Close()
}

func salvaStatus(db *sql.DB, id int) {
	stmt, err := db.Prepare("update restricoes set processado = 'S' where id = ?")
	util.CheckErr(err)
	_, err = stmt.Exec(id)
	util.CheckErr(err)
	stmt.Close()
}
