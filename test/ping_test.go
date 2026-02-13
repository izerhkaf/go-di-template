package test

// func testPing(t *testing.T) {
// 	routes := fmt.Sprintf("%sping", Version)
// 	req := httptest.NewRequest("GET", routes, nil)

// 	resp := httptest.NewRecorder()
// 	App.ServeHTTP(resp, req)

// 	bodyBytes := resp.Body.Bytes()

// 	var body Response[dto.PingResponse]
// 	err := json.Unmarshal(bodyBytes, &body)
// 	require.NoError(t, err)

// 	require.Equal(t, 200, resp.Code)
// 	require.Equal(t, "pong", body.Message)
// }
