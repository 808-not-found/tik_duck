package controller

// nolint:all
var DemoVideos = []Video{
	{
		ID:            1,
		Author:        DemoUser,
		PlayURL:       "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

// nolint:all
var DemoComments = []Comment{
	{
		ID:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

// nolint:all
var DemoUser = User{
	ID:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
