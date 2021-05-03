package repository

import (
	"database/sql"
	"fmt"
	"gallery-backend/model"
	. "gallery-backend/utils"
)

func GetExhibitionList() (exhibitionList []model.Exhibition, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			exhibitionList = nil
			ErrorLog.Println(panic_err)
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	queryStr := "SELECT id, title, description FROM exhibition ORDER BY id"

	rows, err := galleryDB.Query(queryStr)
	if err != nil {
		return nil, err
	}

	exhibitionList = make([]model.Exhibition, 0)
	for rows.Next() {
		var id int64
		var title, description string

		if err := rows.Scan(&id, &title, &description); err != nil {
			return nil, err
		}

		exhibitionList = append(exhibitionList, model.Exhibition{
			ID: id, Title: title, Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exhibitionList, nil
}

func GetExhibitionPaintInfoByExhibitionID(exhibitionID int64) (paintList []model.Paint, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			paintList = nil
			ErrorLog.Println(err)
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	queryStr := "SELECT `paint`.`id`, `paint`.`name`, `paint`.`image_url` FROM `exhibition` " +
		"INNER JOIN `exhibition_has_paint` ON `exhibition`.`id` = `exhibition_has_paint`.`exibition_id` " +
		"INNER JOIN `paint` ON `exhibition_has_paint`.`paint_id` = `paint`.`id` " +
		"WHERE `exhibition`.`id` = ? ORDER BY `paint`.`id`"

	stmt, err := galleryDB.Prepare(queryStr)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(exhibitionID)
	if err != nil {
		return nil, err
	}

	var id int64
	var name, url string
	paintList = make([]model.Paint, 0)

	for rows.Next() {
		err := rows.Scan(&id, &name, &url)
		if err != nil {
			return nil, err
		}

		paintList = append(paintList, model.Paint{ID: id, Name: name, Url: url})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paintList, nil
}

func CreateNewExhibition(title, description string) (lastInsertID int64, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			ErrorLog.Println(panic_err)
			lastInsertID = -1
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	prepareStr := "INSERT INTO exhibition (title, description) VALUES (?, ?)"

	stmt, err := galleryDB.Prepare(prepareStr)
	if err != nil {
		ErrorLog.Println(err)
		return -1, err
	}

	result, err := stmt.Exec(title, description)
	if err != nil {
		ErrorLog.Println(err)
		return -1, err
	}

	lastInsertID, err = result.LastInsertId()
	if err != nil {
		ErrorLog.Println(err)
		return -1, err
	}

	return lastInsertID, nil
}

func AddPaintToExhibition(exhibitionID int64, paintIDList []int64) (err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			ErrorLog.Println(panic_err)
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	prepareStr := "REPLACE INTO exhibition_has_paint( exibition_id, paint_id) VALUES (?, ?)"

	stmt, err := galleryDB.Prepare(prepareStr)
	if err != nil {
		ErrorLog.Println(err)
		return err
	}

	for _, paintID := range paintIDList {
		_, err := stmt.Exec(exhibitionID, paintID)
		if err != nil {
			ErrorLog.Println(err)
		}
	}

	return err
}

func DeletePaintFromExhibition(exhibitionID int64) (err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			ErrorLog.Println(panic_err)
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	prepareStr := "DELETE FROM exhibition_has_paint WHERE exibition_id = ?"

	stmt, err := galleryDB.Prepare(prepareStr)
	if err != nil {
		ErrorLog.Println(err)
		return err
	}

	_, err = stmt.Exec(exhibitionID)
	if err != nil {
		ErrorLog.Println(err)
	}

	return err
}

func GetPaintList() (paintList []model.Paint, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			paintList = nil
			ErrorLog.Println(panic_err)
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	queryStr := "SELECT id, name, image_url FROM paint ORDER BY id"

	rows, err := galleryDB.Query(queryStr)
	if err != nil {
		return nil, err
	}

	var id int64
	var name, url string
	var imageUrl sql.NullString
	paintList = make([]model.Paint, 0)

	for rows.Next() {
		err := rows.Scan(&id, &name, &imageUrl)
		if err != nil {
			return nil, err
		}

		if imageUrl.Valid {
			url = imageUrl.String
		} else {
			url = ""
		}

		paintList = append(paintList, model.Paint{ID: id, Name: name, Url: url})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paintList, nil
}

func CreateNewPaint(paintName string) (lastInsertID int64, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			ErrorLog.Println(panic_err)
			lastInsertID = -1
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	prepareStr := "INSERT INTO paint (name) VALUES (?)"

	stmt, err := galleryDB.Prepare(prepareStr)
	if err != nil {
		ErrorLog.Println(err)
		return -1, err
	}

	result, err := stmt.Exec(paintName)
	if err != nil {
		ErrorLog.Println(err)
		return -1, err
	}

	lastInsertID, err = result.LastInsertId()
	if err != nil {
		ErrorLog.Println(err)
		return -1, err
	}

	return lastInsertID, nil
}

func UpdatePaintImageUrl(paintID int64, paintUrl string) (err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if panic_err := recover(); panic_err != nil {
			ErrorLog.Println(panic_err)
			err = fmt.Errorf("%v", panic_err)
		}

		return
	}()

	prepareStr := "UPDATE paint SET image_url = ? WHERE id = ? "

	stmt, err := galleryDB.Prepare(prepareStr)
	if err != nil {
		ErrorLog.Println(err)
		return err
	}

	_, err = stmt.Exec(paintUrl, paintID)
	if err != nil {
		ErrorLog.Println(err)
		return err
	}

	return nil
}
