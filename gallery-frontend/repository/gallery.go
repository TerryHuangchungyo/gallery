package repository

import (
	"gallery-frontend/model"
	. "gallery-frontend/utils"
)

func GetExhibitionList() (exhibitionList []map[string]interface{}, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if err := recover(); err != nil {
			exhibitionList = nil
			ErrorLog.Println(err)
		}

		return
	}()

	queryStr := "SELECT id, title FROM exhibition"

	rows, err := galleryDB.Query(queryStr)
	if err != nil {
		return nil, err
	}

	exhibitionList = make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int64
		var title string

		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}

		data := make(map[string]interface{})
		data["id"] = id
		data["title"] = title

		exhibitionList = append(exhibitionList, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exhibitionList, nil
}

func GetExhibitionInfoByExhibitionID(exhibitionID int64) (exhibition *model.Exhibition, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if err := recover(); err != nil {
			exhibition = nil
			ErrorLog.Println(err)
		}

		return
	}()

	queryStr := "SELECT id, title, description FROM `exhibition` WHERE id = ?"

	stmt, err := galleryDB.Prepare(queryStr)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(exhibitionID)
	if err != nil {
		return nil, err
	}

	var id int64
	var title, description string
	exhibition = &model.Exhibition{}

	for rows.Next() {
		err := rows.Scan(&id, &title, &description)
		if err != nil {
			return nil, err
		}

		exhibition.ID = id
		exhibition.Title = title
		exhibition.Description = description
		exhibition.Paints = make([]model.Paint, 0)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exhibition, nil
}

func GetExhibitionPaintInfoByExhibitionID(exhibitionID int64) (paintList []model.Paint, err error) {
	defer func() {
		if err != nil {
			ErrorLog.Println(err)
		}

		if err := recover(); err != nil {
			paintList = nil
			ErrorLog.Println(err)
		}

		return
	}()

	queryStr := "SELECT `paint`.`id`, `paint`.`name`, `paint`.`image_url` FROM `exhibition` " +
		"INNER JOIN `exhibition_has_paint` ON `exhibition`.`id` = `exhibition_has_paint`.`exibition_id` " +
		"INNER JOIN `paint` ON `exhibition_has_paint`.`paint_id` = `paint`.`id` " +
		"WHERE `exhibition`.`id` = ?"

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
