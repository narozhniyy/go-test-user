package connLogRepository

import (
	"github.com/narozhniyy/test/models"
	"github.com/narozhniyy/test/resources"
	"log"
)

type connLogRepository struct {
}

func InitConnLogRepository() *connLogRepository {
	return &connLogRepository{}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (r *connLogRepository) Compare(fu int64, su int64) ([]models.IpOccurrence, error) {
	var item models.IpOccurrence
	var ipOccurrences []models.IpOccurrence
	rows, err := resources.GetDB().Query(`
			SELECT count(ip_addr) numb 
			FROM (
			  SELECT distinct ip_addr 
			  FROM conn_log 
			  WHERE user_id = ? 
				
			  UNION ALL
			  
			  SELECT distinct ip_addr 
			  FROM conn_log 
			  WHERE user_id = ?
			) AS sub 
			GROUP BY ip_addr 
			HAVING numb >= 2`,
	fu, su)

	defer func() {
		if err := rows.Close(); err != nil {
			logFatal(err)
		}
	}()

	logFatal(err)

	for rows.Next() {
		err := rows.Scan(&item.Numb)
		logFatal(err)
		ipOccurrences = append(ipOccurrences, item)
	}

	return ipOccurrences, err
}