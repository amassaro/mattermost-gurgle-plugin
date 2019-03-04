package main

// func TestServeHTTP(t *testing.T) {
// 	assert := assert.New(t)
// 	plugin := GurglePlugin{}
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/", nil)

// 	plugin.ServeHTTP(nil, w, r)

// 	result := w.Result()
// 	assert.NotNil(result)
// 	bodyBytes, err := ioutil.ReadAll(result.Body)
// 	assert.Nil(err)
// 	bodyString := string(bodyBytes)

// 	// assert.Equal("~"+plugin.configuration.ChannelName+" - @"+plugin.configuration.TeamName, bodyString)
// }
