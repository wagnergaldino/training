package geocode

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/codingsince1985/geo-golang/bing"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/wagnergaldino/training/dne/models"
	"github.com/wagnergaldino/training/dne/util"
)

var keys = map[int]string{
	0: "AtNk7Iwo3YwbVHZ6MDgEhUCtmuf9IsUuoaBIzuVA49Nlii_krwSSrWdQpsetdY9T",
	1: "AuktwWF5Je9QtdAbLN_0t0jdPsL2KFRa1Mo_1Ny-46PjdacHWmQbejvo-AmyWCY6",
	2: "AkONyq1Qla5_gOxo84EJE0XZTK2i3Z0ocriAgimf-OLKiwd-rmLs1ye2CKQ4E16u",
	3: "As4BqbhW1codxcH0EYpyWa9AOgRaVPCZHgWy7zS_yXaskZ4sqTA3i8c3PCiVgQjK"}

func Parallel(db *gorm.DB) {

	for {

		var addresses []models.Address

		err := db.Model(models.Address{}).Where("latitude is null").Limit(1000).Find(&addresses).Error

		if err != nil {
			log.Println(err)
		}
		spew.Dump(err)

		if len(addresses) == 0 {
			break
		}

		// set the number of channels
		channels := len(addresses)

		// set the number of simultaneous addresses
		simultaneous := 4

		// there is less addresses to run than parallel argument
		if channels < simultaneous {
			simultaneous = channels
		}

		// init channels
		queue := make(chan models.Address, channels)

		// init wait group
		wg := &sync.WaitGroup{}

		// generate tasks
		for _, process := range addresses {
			queue <- process
		}

		// start workers
		for i := 0; i < simultaneous; i++ {

			// add worker
			wg.Add(1)

			// go concurrency
			go func(wg *sync.WaitGroup) {

				// make concurrency safe (finish worker)
				defer wg.Done()

				// call process
				for q := range queue {

					// check if address was successfully processed
					if err = Process(&q, db, keys[i]); err != nil {
						log.Println(err)
					}

				}

			}(wg) // this is a closure, wg must be passed as a pointer

		}

		// close queue of tasks
		close(queue)

		// wait for all the workers to finish
		wg.Wait()

	}

}

func Process(address *models.Address, db *gorm.DB, key string) error {
	query := fmt.Sprintf("%s, %s, %s, %s, BR, %s", address.Logradouro, address.Bairro, address.Localidade, address.Uf, address.Cep)

	address.Latitude, address.Longitude = getGeocode(query, key)

	return db.Save(address).Error
}

func salvaCep(db *sql.DB, latitude, longitude, cep string) {
	stmt, err := db.Prepare("update address set latitude = ?, longitude = ? where cep = ?")
	stmt.Close()
	util.CheckErr(err)

	_, err = stmt.Exec(latitude, longitude, cep)
	util.CheckErr(err)
}

func getGeocode(address, key string) (latitude, longitude string) {
	geocoder := bing.Geocoder(key)
	location, err := geocoder.Geocode(address)
	if err != nil {
		fmt.Println(err)
		latitude = ""
		longitude = ""
	} else {
		latitude = fmt.Sprint(location.Lat)
		longitude = fmt.Sprint(location.Lng)
	}
	return
}
