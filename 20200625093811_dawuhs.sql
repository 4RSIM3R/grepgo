-- +goose Up
CREATE TABLE IF NOT EXISTS dawuhs(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Dawuh VARCHAR (255) ,
    Qoil VARCHAR (255) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
);

-- +goose Down
DROP TABLE dawuhs;
