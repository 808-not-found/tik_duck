package db

import(
	  "time"
	  
      "github.com/808-not-found/tik_duck/pkg/consts"
	  "gorm.io/gorm"
)

type Like struct{
	 gorm.Model
	 ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	 LikeTime    time.Time `gorm:"column:like_time;default:CURRENT_TIMESTAMP;NOT NULL"`
     UserID        int64      `gorm:"column:user_id;default:0;NOT NULL"`
     VideoID       int64      `gorm:"column:video_id;default:0;NOT NULL"`

}

func(like *Like)TableName() string {
	return consts.LikeTableName
}



