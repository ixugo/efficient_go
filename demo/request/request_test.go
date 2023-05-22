package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := http.Get("http://asd.gr.com/123/124/123")
	fmt.Printf("vv: %#v\n", err)

	if err, ok := err.(*url.Error); ok {
		fmt.Println(ok)
		fmt.Println(err.Err)
	}
	err = errors.Unwrap(err)
	fmt.Printf("v: %+v\n", err)

	fmt.Println(errors.Unwrap(nil))
	_ = resp

}

type User struct {
	ID int `json:"id,string"`
}

func TestTime(t *testing.T) {
	var u User
	u.ID = 123
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	const str = `{"id":"123"}`

	var newUser User
	err := json.Unmarshal([]byte(str), &newUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(newUser.ID)
}

const str = `{
	"Ability" :
	[
		{
			"code" : 8,
			"desc" : "\u660e\u70df\u660e\u706b\u68c0\u6d4b",
			"item" : 46,
			"name" : "\u660e\u70df\u660e\u706b\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u68c0\u6d4b\u5230\u660e\u706b",
					"property" : "Fire"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u660e\u70df",
					"property" : "Smoke"
				}
			],
			"sub" : true
		},
		{
			"code" : 62,
			"desc" : "\u533a\u57df\u8f66\u8f86\u7981\u505c",
			"item" : 48,
			"name" : "\u533a\u57df\u8f66\u8f86\u7981\u505c",
			"parameters" :
			[
				{
					"default" : "5",
					"key" : "parking_minute",
					"max" : "",
					"min" : "",
					"name" : "\u505c\u9760\u65f6\u957f(\u5206)",
					"required" : true,
					"type" : 2
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u8f66\u8f86\u957f\u65f6\u95f4\u505c\u6cca",
					"property" : "NoParking"
				}
			],
			"sub" : true
		},
		{
			"code" : 62,
			"desc" : "\u8f66\u724c\u8bc6\u522b",
			"item" : 63,
			"name" : "\u8f66\u724c\u8bc6\u522b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u8f66\u724c",
					"property" : "Plate"
				}
			],
			"sub" : true
		},
		{
			"code" : 59,
			"desc" : "\u5c0f\u52a8\u7269\u68c0\u6d4b",
			"item" : 59,
			"name" : "\u5c0f\u52a8\u7269\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u68c0\u6d4b\u5230\u5c0f\u52a8\u7269",
					"property" : "Animal"
				}
			],
			"sub" : false
		},
		{
			"code" : 54,
			"desc" : "\u4eba\u5458\u62e5\u6324\u68c0\u6d4b",
			"item" : 54,
			"name" : "\u4eba\u5458\u62e5\u6324\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u62e5\u6324\u5ea6-\u62e5\u6324",
					"property" : "Crowd-Busy"
				},
				{
					"name" : "\u62e5\u6324\u5ea6-\u6b63\u5e38",
					"property" : "Crowd-Normal"
				},
				{
					"name" : "\u62e5\u6324\u5ea6-\u7a7a\u65f7",
					"property" : "Crowd-Free"
				}
			],
			"sub" : false
		},
		{
			"code" : 58,
			"desc" : "\u53e3\u7f69\u68c0\u6d4b",
			"item" : 246,
			"name" : "\u53e3\u7f69\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u672a\u4f69\u6234\u53e3\u7f69",
					"property" : "NoMask"
				},
				{
					"name" : "\u53e3\u7f69\u4f69\u6234\u4e0d\u89c4\u8303",
					"property" : "BadMask"
				}
			],
			"sub" : true
		},
		{
			"code" : 58,
			"desc" : "\u8138\u90e8\u6293\u62cd",
			"item" : 248,
			"name" : "\u8138\u90e8\u6293\u62cd",
			"parameters" :
			[
				{
					"default" : "0.65",
					"key" : "front_face_threshold",
					"max" : "",
					"min" : "",
					"name" : "\u6293\u62cd\u9608\u503c",
					"required" : false,
					"type" : 2
				},
				{
					"default" : "40",
					"key" : "front_face_min_pixel",
					"max" : "",
					"min" : "",
					"name" : "\u6700\u5c0f\u4eba\u8138\u50cf\u7d20",
					"required" : false,
					"type" : 0
				},
				{
					"default" : "500",
					"key" : "face_repeat_upload_ms",
					"max" : "",
					"min" : "",
					"name" : "\u66f4\u65b0\u95f4\u9694(\u8c6a\u79d2)",
					"required" : false,
					"type" : 2
				},
				{
					"default" : "false",
					"key" : "expand_face_area",
					"max" : "",
					"min" : "",
					"name" : "\u6269\u5927\u8138\u90e8\u533a\u57df",
					"required" : false,
					"type" : 4
				},
				{
					"default" : "true",
					"key" : "upload_cropped_face_image",
					"max" : "",
					"min" : "",
					"name" : "\u4e0a\u4f20\u88c1\u526a\u7684\u4eba\u8138\u56fe",
					"required" : false,
					"type" : 4
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u6293\u62cd\u5230\u4eba\u8138",
					"property" : "CaptureFace"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u4eba\u6570\u8d85\u5458\u68c0\u6d4b",
			"item" : 57,
			"name" : "\u4eba\u6570\u8d85\u5458\u68c0\u6d4b",
			"parameters" :
			[
				{
					"default" : "5",
					"key" : "limit_number",
					"max" : "",
					"min" : "",
					"name" : "\u4eba\u6570\u4e0a\u9650",
					"required" : true,
					"type" : 0
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u5de5\u4f5c\u533a\u8d85\u5458",
					"property" : "ZoneOverload"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u4eba\u6570\u76d1\u63a7",
			"item" : 230,
			"name" : "\u4eba\u6570\u76d1\u63a7",
			"parameters" :
			[
				{
					"default" : "true",
					"key" : "upload_on_people_num_changed",
					"max" : "",
					"min" : "",
					"name" : "\u4ec5\u4eba\u6570\u53d8\u5316\u65f6\u4e0a\u62a5",
					"required" : false,
					"type" : 4
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u753b\u9762\u4eba\u6570\u4e0a\u62a5",
					"property" : "CurrentPeopleNum"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u79bb\u5c97\u68c0\u6d4b",
			"item" : 7,
			"name" : "\u79bb\u5c97\u68c0\u6d4b",
			"parameters" :
			[
				{
					"default" : "180",
					"key" : "staff_sec",
					"max" : "",
					"min" : "",
					"name" : "\u8d85\u65f6\u65f6\u95f4(\u79d2)",
					"required" : true,
					"type" : 0
				},
				{
					"default" : "1",
					"key" : "staff_number",
					"max" : "",
					"min" : "",
					"name" : "\u8981\u6c42\u5728\u5c97\u4eba\u6570",
					"required" : true,
					"type" : 0
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u5728\u5c97\u4eba\u5458\u4e0d\u8db3",
					"property" : "StaffAway"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u8d8a\u7ebf\u68c0\u6d4b",
			"item" : 16,
			"name" : "\u8d8a\u7ebf\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u4eba\u5458\u8d8a\u7ebf",
					"property" : "PeopleCross"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u533a\u57df\u5165\u4fb5",
			"item" : 9,
			"name" : "\u533a\u57df\u5165\u4fb5",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u4eba\u5458\u95ef\u5165",
					"property" : "Intrude"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u4eba\u5458\u5f98\u5f8a",
			"item" : 228,
			"name" : "\u4eba\u5458\u5f98\u5f8a",
			"parameters" :
			[
				{
					"default" : "60.0",
					"key" : "people_wander_sec",
					"max" : "",
					"min" : "",
					"name" : "\u6700\u5927\u5f98\u5f8a\u65f6\u957f(\u79d2)",
					"required" : false,
					"type" : 2
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u4eba\u5458\u5f98\u5f8a",
					"property" : "Wandering"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u7ffb\u8d8a\u56f4\u680f",
			"item" : 10,
			"name" : "\u7ffb\u8d8a\u56f4\u680f",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u7ffb\u8d8a\u56f4\u680f",
					"property" : "TouchFence"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u9759\u6b62\u68c0\u6d4b",
			"item" : 226,
			"name" : "\u9759\u6b62\u68c0\u6d4b",
			"parameters" :
			[
				{
					"default" : "30.0",
					"key" : "people_stationary_sec",
					"max" : "",
					"min" : "",
					"name" : "\u6700\u5927\u9759\u6b62\u65f6\u957f(\u79d2)",
					"required" : false,
					"type" : 2
				}
			],
			"permitted" : false,
			"policy" :
			[
				{
					"name" : "\u957f\u65f6\u95f4\u9759\u6b62",
					"property" : "Stationary"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u4eba\u5458\u805a\u96c6",
			"item" : 224,
			"name" : "\u4eba\u5458\u805a\u96c6",
			"parameters" :
			[
				{
					"default" : "3",
					"key" : "cluster_people_threshold",
					"max" : "",
					"min" : "",
					"name" : "\u4eba\u6570\u8bbe\u7f6e",
					"required" : false,
					"type" : 0
				},
				{
					"default" : "1.2",
					"key" : "cluster_width_expand_ratio",
					"max" : "",
					"min" : "",
					"name" : "\u805a\u96c6\u7cfb\u6570(\u8d8a\u5927\u8d8a\u5bbd\u677e)",
					"required" : false,
					"type" : 2
				}
			],
			"permitted" : false,
			"policy" :
			[
				{
					"name" : "\u4eba\u5458\u805a\u96c6",
					"property" : "PeopleGather"
				}
			],
			"sub" : true
		},
		{
			"code" : 61,
			"desc" : "\u6500\u722c\u68c0\u6d4b",
			"item" : 61,
			"name" : "\u6500\u722c\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u68c0\u6d4b\u5230\u6500\u722c",
					"property" : "Climb"
				}
			],
			"sub" : false
		},
		{
			"code" : 15,
			"desc" : "\u5ba2\u6d41\u8ba1\u6570",
			"item" : 15,
			"name" : "\u5ba2\u6d41\u8ba1\u6570",
			"parameters" :
			[
				{
					"default" : "false",
					"key" : "upload_image_on_count_people",
					"max" : "",
					"min" : "",
					"name" : "\u4e0a\u62a5\u65f6\u662f\u5426\u4e0a\u4f20\u56fe\u7247",
					"required" : false,
					"type" : 4
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u4eba\u5458\u8fdb\u51fa",
					"property" : "HeadCount"
				}
			],
			"sub" : false
		},
		{
			"code" : 14,
			"desc" : "\u6253\u67b6\u68c0\u6d4b",
			"item" : 14,
			"name" : "\u6253\u67b6\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u7591\u4f3c\u6253\u67b6",
					"property" : "Fighting"
				}
			],
			"sub" : false
		},
		{
			"code" : 13,
			"desc" : "\u7535\u52a8\u8f66\u68c0\u6d4b",
			"item" : 13,
			"name" : "\u7535\u52a8\u8f66\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u68c0\u6d4b\u5230\u7535\u52a8\u8f66",
					"property" : "EBike"
				}
			],
			"sub" : false
		},
		{
			"code" : 55,
			"desc" : "\u5360\u9053\u7ecf\u8425\u68c0\u6d4b",
			"item" : 55,
			"name" : "\u5360\u9053\u7ecf\u8425\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u68c0\u6d4b\u5230\u644a\u4f4d",
					"property" : "Stall"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u906e\u9633\u7bf7",
					"property" : "Sunshade"
				}
			],
			"sub" : false
		},
		{
			"code" : 56,
			"desc" : "\u8857\u9053\u5783\u573e\u68c0\u6d4b",
			"item" : 56,
			"name" : "\u8857\u9053\u5783\u573e\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u68c0\u6d4b\u5230\u70df\u5934",
					"property" : "Cigarette"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u5efa\u7b51\u5783\u573e",
					"property" : "Construction"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u5783\u573e\u74f6",
					"property" : "Bottle"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u7eb8\u7bb1",
					"property" : "Carton"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u7eb8\u5f20\u5783\u573e",
					"property" : "Paper"
				},
				{
					"name" : "\u68c0\u6d4b\u5230\u5783\u573e\u888b",
					"property" : "Bag"
				}
			],
			"sub" : false
		},
		{
			"code" : 1,
			"desc" : "\u5b89\u5168\u5e3d\u68c0\u6d4b",
			"item" : 2,
			"name" : "\u5b89\u5168\u5e3d\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u672a\u6234\u5b89\u5168\u5e3d",
					"property" : "NoHelmet"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u53cd\u5149\u8863\u68c0\u6d4b",
			"item" : 5,
			"name" : "\u53cd\u5149\u8863\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u672a\u7a7f\u53cd\u5149\u8863",
					"property" : "NoVest"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u672a\u7a7f\u957f\u8896\u68c0\u6d4b",
			"item" : 6,
			"name" : "\u672a\u7a7f\u957f\u8896\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u672a\u7a7f\u957f\u8896",
					"property" : "NoSleeve"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u5de5\u670d\u68c0\u6d4b",
			"item" : 60,
			"name" : "\u5de5\u670d\u68c0\u6d4b",
			"parameters" :
			[
				{
					"default" : "true",
					"key" : "require_all_selected_suite",
					"max" : "",
					"min" : "",
					"name" : "\u8981\u6c42\u5339\u914d\u6240\u6709\u52fe\u9009\u6a21\u677f",
					"required" : false,
					"type" : 4
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u672a\u7a7f\u5de5\u670d",
					"property" : "NoSuit"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u7761\u5c97\u68c0\u6d4b",
			"item" : 11,
			"name" : "\u7761\u5c97\u68c0\u6d4b",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u7591\u4f3c\u7761\u5c97",
					"property" : "Sleeping"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u5de5\u5730\u5b89\u5168\u5e26",
			"item" : 50,
			"name" : "\u5de5\u5730\u5b89\u5168\u5e26",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u65e0\u5b89\u5168\u5e26",
					"property" : "NoBelt"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u4eba\u5458\u5012\u5730",
			"item" : 3,
			"name" : "\u4eba\u5458\u5012\u5730",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u4eba\u5458\u5012\u5730",
					"property" : "Falldown"
				}
			],
			"sub" : true
		},
		{
			"code" : 1,
			"desc" : "\u62bd\u70df\u6253\u7535\u8bdd\u68c0\u6d4b",
			"item" : 4,
			"name" : "\u62bd\u70df\u6253\u7535\u8bdd\u68c0\u6d4b",
			"parameters" :
			[
				{
					"default" : "true",
					"key" : "behavior_enable_calling",
					"max" : "",
					"min" : "",
					"name" : "\u5f00\u542f\u6253\u7535\u8bdd\u68c0\u6d4b",
					"required" : false,
					"type" : 4
				},
				{
					"default" : "true",
					"key" : "behavior_enable_smoking",
					"max" : "",
					"min" : "",
					"name" : "\u5f00\u542f\u62bd\u70df\u68c0\u6d4b",
					"required" : false,
					"type" : 4
				},
				{
					"default" : "true",
					"key" : "behavior_enable_playing",
					"max" : "",
					"min" : "",
					"name" : "\u5f00\u542f\u73a9\u624b\u673a\u68c0\u6d4b",
					"required" : false,
					"type" : 4
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u62bd\u70df",
					"property" : "Smoking"
				},
				{
					"name" : "\u73a9\u624b\u673a",
					"property" : "Playing"
				},
				{
					"name" : "\u6253\u7535\u8bdd",
					"property" : "Calling"
				}
			],
			"sub" : true
		},
		{
			"code" : 52,
			"desc" : "\u9a7e\u9a76\u5458\u5de6\u987e\u53f3\u76fc",
			"item" : 51,
			"name" : "\u9a7e\u9a76\u5458\u5de6\u987e\u53f3\u76fc",
			"parameters" :
			[
				{
					"default" : "3",
					"key" : "driver_careless_sec",
					"max" : "",
					"min" : "",
					"name" : "\u8d85\u65f6\u65f6\u95f4(\u79d2)",
					"required" : true,
					"type" : 0
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u9a7e\u9a76\u5458\u5de6\u987e\u53f3\u76fc",
					"property" : "DriverCareless"
				}
			],
			"sub" : true
		},
		{
			"code" : 49,
			"desc" : "\u65e0\u6d88\u9632\u5668\u6750",
			"item" : 49,
			"name" : "\u65e0\u6d88\u9632\u5668\u6750",
			"parameters" :
			[
				{
					"default" : "5",
					"key" : "extinguisher_disappear_second",
					"max" : "",
					"min" : "",
					"name" : "\u8d85\u65f6\u65f6\u95f4(\u79d2)",
					"required" : true,
					"type" : 0
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u672a\u653e\u7f6e\u6d88\u9632\u5668\u6750",
					"property" : "NoExtinguisher"
				}
			],
			"sub" : false
		},
		{
			"code" : 47,
			"desc" : "\u5de5\u7a0b\u5668\u68b0",
			"item" : 239,
			"name" : "\u5de5\u7a0b\u5668\u68b0",
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u540a\u8f66",
					"property" : "Machine-DiaoChe"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u5854\u540a",
					"property" : "Machine-TaDiao"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u6c34\u6ce5\u6cf5\u8f66",
					"property" : "Machine-ShuiNiBengChe"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u63a8\u571f\u673a",
					"property" : "Machine-TuiTuJi"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u94f2\u8f66",
					"property" : "Machine-ChanChe"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u6316\u6398\u673a",
					"property" : "Machine-WaJueJi"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u7ffb\u6597\u8f66",
					"property" : "Machine-FanDouChe"
				},
				{
					"name" : "\u5de5\u7a0b\u5668\u68b0-\u5176\u4ed6\u5668\u68b0",
					"property" : "Machine-QiTaShiGongJiXie"
				}
			],
			"sub" : true
		},
		{
			"code" : 47,
			"desc" : "\u540a\u81c2\u4e0b\u7ad9\u4eba",
			"item" : 240,
			"name" : "\u540a\u81c2\u4e0b\u7ad9\u4eba",
			"permitted" : false,
			"policy" :
			[
				{
					"name" : "\u540a\u81c2\u4e0b\u7ad9\u4eba",
					"property" : "PeopleUnderMachine"
				}
			],
			"sub" : true
		},
		{
			"code" : 255,
			"desc" : "\u975e\u673a\u52a8\u8f66\u505c\u653e",
			"item" : 255,
			"name" : "\u975e\u673a\u52a8\u8f66\u505c\u653e",
			"parameters" :
			[
				{
					"default" : "1",
					"key" : "fjdc_parking_minute",
					"max" : "",
					"min" : "",
					"name" : "\u8d85\u65f6\u65f6\u95f4(\u5206\u949f)",
					"required" : true,
					"type" : 0
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u975e\u673a\u52a8\u8f66\u957f\u65f6\u95f4\u505c\u653e",
					"property" : "FJDC-Parking"
				}
			],
			"sub" : false
		},
		{
			"code" : 40,
			"desc" : "\u901a\u9053\u5360\u7528",
			"item" : 40,
			"name" : "\u901a\u9053\u5360\u7528",
			"parameters" :
			[
				{
					"default" : "3",
					"key" : "chn_occupy_minute",
					"max" : "",
					"min" : "",
					"name" : "\u8d85\u65f6\u65f6\u95f4(\u5206\u949f)",
					"required" : true,
					"type" : 0
				},
				{
					"default" : "5",
					"key" : "chn_occupy_update_minute",
					"max" : "",
					"min" : "",
					"name" : "\u80cc\u666f\u66f4\u65b0\u95f4\u9694(\u5206\u949f)",
					"required" : true,
					"type" : 0
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u901a\u9053\u5360\u7528",
					"property" : "Occupied"
				}
			],
			"sub" : false
		},
		{
			"code" : 45,
			"desc" : "\u9759\u7535\u91ca\u653e",
			"item" : 229,
			"name" : "\u9759\u7535\u91ca\u653e",
			"parameters" :
			[
				{
					"default" : "5",
					"key" : "esd_time_threshold",
					"max" : "60",
					"min" : "1",
					"name" : "\u6700\u4f4e\u91ca\u653e\u65f6\u95f4",
					"required" : true,
					"type" : 2
				}
			],
			"permitted" : false,
			"policy" :
			[
				{
					"name" : "\u9759\u7535\u91ca\u653e\u65f6\u95f4\u4e0d\u8db3",
					"property" : "BadDischarge"
				}
			],
			"sub" : true
		},
		{
			"code" : 23,
			"desc" : "\u4eba\u8138\u8bc6\u522b",
			"item" : 24,
			"name" : "\u4eba\u8138\u8bc6\u522b",
			"parameters" :
			[
				{
					"default" : "0.75",
					"key" : "face_reg_similarity",
					"max" : "1",
					"min" : "0.7",
					"name" : "\u76f8\u4f3c\u5ea6\u9608\u503c",
					"required" : true,
					"type" : 2
				},
				{
					"default" : "0.3",
					"key" : "face_low_similarity",
					"max" : "0.6",
					"min" : "0.01",
					"name" : "\u964c\u751f\u4eba\u9608\u503c",
					"required" : true,
					"type" : 2
				}
			],
			"permitted" : true,
			"policy" :
			[
				{
					"name" : "\u8bc6\u522b\u5230\u5728\u518c\u4eba\u8138",
					"property" : "FaceId"
				},
				{
					"name" : "\u8bc6\u522b\u5230\u5176\u4ed6\u4eba\u8138",
					"property" : "Stranger"
				}
			],
			"sub" : true
		}
	],
	"BoardId" : "RJ-BMOX-CD46180F3A9D08097CD8EC8725802F7F",
	"BoardIp" : "192.168.1.240",
	"Event" : "/alg_ability_fetch",
	"Result" :
	{
		"Code" : 0,
		"Desc" : "Success"
	}
}
`

type AutoGenerated struct {
	Ability []struct {
		Code      int    `json:"code"`
		Desc      string `json:"desc"`
		Item      int    `json:"item"`
		Name      string `json:"name"`
		Permitted bool   `json:"permitted"`
		Policy    []struct {
			Name     string `json:"name"`
			Property string `json:"property"`
		} `json:"policy"`
		Sub        bool `json:"sub"`
		Parameters []struct {
			Default  string `json:"default"`
			Key      string `json:"key"`
			Max      string `json:"max"`
			Min      string `json:"min"`
			Name     string `json:"name"`
			Required bool   `json:"required"`
			Type     int    `json:"type"`
		} `json:"parameters,omitempty"`
	} `json:"Ability"`
	BoardID string `json:"BoardId"`
	BoardIP string `json:"BoardIp"`
	Event   string `json:"Event"`
	Result  struct {
		Code int    `json:"Code"`
		Desc string `json:"Desc"`
	} `json:"Result"`
}

func TestJSON(t *testing.T) {
	var data AutoGenerated
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		t.Fatal(err)
	}

	for _, vs := range data.Ability {
		for _, v := range vs.Policy {
			fmt.Println(v.Property, v.Name)
		}
	}
}
