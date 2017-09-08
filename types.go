package main

// TopPosts holds the Posts array
type TopPosts struct {
	Posts []Post
}

// Post type that retrieves the basic post information
type Post struct {
	ID         string `json:"id"`
	Dimensions struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"dimensions"`
	Owner struct {
		ID string `json:"id"`
	} `json:"owner"`
	ThumbnailSrc       string        `json:"thumbnail_src"`
	ThumbnailResources []interface{} `json:"thumbnail_resources"`
	IsVideo            bool          `json:"is_video"`
	Code               string        `json:"code"`
	Date               int           `json:"date"`
	DisplaySrc         string        `json:"display_src"`
	Caption            string        `json:"caption"`
	Comments           struct {
		Count int `json:"count"`
	} `json:"comments"`
	Likes struct {
		Count int `json:"count"`
	} `json:"likes"`
	VideoViews int `json:"video_views,omitempty"`
}

// Instagram type that decodes JSON output
type Instagram struct {
	ActivityCounts interface{} `json:"activity_counts"`
	Config         struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
	} `json:"config"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	EntryData    struct {
		LocationsPage []struct {
			Location struct {
				ID            string  `json:"id"`
				Name          string  `json:"name"`
				HasPublicPage bool    `json:"has_public_page"`
				Lat           float64 `json:"lat"`
				Lng           float64 `json:"lng"`
				Slug          string  `json:"slug"`
				Media         struct {
					Nodes []struct {
						CommentsDisabled bool   `json:"comments_disabled"`
						ID               string `json:"id"`
						Dimensions       struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						Owner struct {
							ID string `json:"id"`
						} `json:"owner"`
						ThumbnailSrc       string        `json:"thumbnail_src"`
						ThumbnailResources []interface{} `json:"thumbnail_resources"`
						IsVideo            bool          `json:"is_video"`
						Code               string        `json:"code"`
						Date               int           `json:"date"`
						DisplaySrc         string        `json:"display_src"`
						Caption            string        `json:"caption,omitempty"`
						Comments           struct {
							Count int `json:"count"`
						} `json:"comments"`
						Likes struct {
							Count int `json:"count"`
						} `json:"likes"`
						VideoViews int `json:"video_views,omitempty"`
					} `json:"nodes"`
					Count    int `json:"count"`
					PageInfo struct {
						HasNextPage bool   `json:"has_next_page"`
						EndCursor   string `json:"end_cursor"`
					} `json:"page_info"`
				} `json:"media"`
				TopPosts struct {
					Nodes []struct {
						ID         string `json:"id"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						Owner struct {
							ID string `json:"id"`
						} `json:"owner"`
						ThumbnailSrc       string        `json:"thumbnail_src"`
						ThumbnailResources []interface{} `json:"thumbnail_resources"`
						IsVideo            bool          `json:"is_video"`
						Code               string        `json:"code"`
						Date               int           `json:"date"`
						DisplaySrc         string        `json:"display_src"`
						Caption            string        `json:"caption"`
						Comments           struct {
							Count int `json:"count"`
						} `json:"comments"`
						Likes struct {
							Count int `json:"count"`
						} `json:"likes"`
						VideoViews int `json:"video_views,omitempty"`
					} `json:"nodes"`
					Count    int `json:"count"`
					PageInfo struct {
						HasNextPage bool        `json:"has_next_page"`
						EndCursor   interface{} `json:"end_cursor"`
					} `json:"page_info"`
				} `json:"top_posts"`
				Directory struct {
					Country struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Slug string `json:"slug"`
					} `json:"country"`
					City struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Slug string `json:"slug"`
					} `json:"city"`
				} `json:"directory"`
			} `json:"location"`
			LoggingPageID string `json:"logging_page_id"`
		} `json:"LocationsPage"`
	} `json:"entry_data"`
	Gatekeepers struct {
		Bn bool `json:"bn"`
		Ld bool `json:"ld"`
		Pl bool `json:"pl"`
	} `json:"gatekeepers"`
	Qe struct {
		Ebd struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ebd"`
		Bc3L struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"bc3l"`
		Ccp struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ccp"`
		CreateUpsell struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"create_upsell"`
		Disc struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"disc"`
		Feed struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"feed"`
		Stories struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stories"`
		SuUniverse struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"su_universe"`
		Us struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"us"`
		UsLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"us_li"`
		Nav struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"nav"`
		NavLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"nav_lo"`
		Profile struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile"`
		Spp struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"spp"`
		Deact struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"deact"`
		Sidecar struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"sidecar"`
		Video struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"video"`
		Filters struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"filters"`
		Typeahead struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"typeahead"`
		LocationTag struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"location_tag"`
		DeltaDefaults struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"delta_defaults"`
		Appsell struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"appsell"`
		FeedSensitivity struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"feed_sensitivity"`
		ProfileSensitivity struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile_sensitivity"`
		Save struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"save"`
		Stale struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stale"`
		Reg struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reg"`
		Nux struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"nux"`
		LikedBy struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"liked_by"`
	} `json:"qe"`
	Hostname                     string `json:"hostname"`
	DisplayPropertiesServerGuess struct {
		PixelRatio    float64 `json:"pixel_ratio"`
		ViewportWidth int     `json:"viewport_width"`
	} `json:"display_properties_server_guess"`
	EnvironmentSwitcherVisibleServerGuess bool   `json:"environment_switcher_visible_server_guess"`
	Platform                              string `json:"platform"`
	ProbablyHasApp                        bool   `json:"probably_has_app"`
	ShowAppInstall                        bool   `json:"show_app_install"`
}
