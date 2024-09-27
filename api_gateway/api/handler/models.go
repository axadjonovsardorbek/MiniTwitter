package handler

type FollowCreateReq struct{
	FollowedId string
}

type LikeCreateReq struct{
	TwitId string
}

type TwitCreateReq struct{
	Content string
	ImageUrl string
}

type ReTwitCreateReq struct{
	Content string
	ImageUrl string
	TwitId string
}