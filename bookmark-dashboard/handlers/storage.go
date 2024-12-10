package handlers

import (
    "bookmark-dashboard/models"
    "encoding/json"
    "io/ioutil"
    "os"
)

var dataFile string

func InitStorage(filePath string) {
    dataFile = filePath
}

func loadData() (*models.Data, error) {
    file, err := os.Open(dataFile)
    if err != nil {
        if os.IsNotExist(err) {
            return &models.Data{Categories: []models.Category{}}, nil
        }
        return nil, err
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var data models.Data
    if err := json.Unmarshal(bytes, &data); err != nil {
        return nil, err
    }
    return &data, nil
}

func saveData(data *models.Data) error {
    bytes, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(dataFile, bytes, 0644)
}
