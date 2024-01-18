CUSTOMERLABS ASSIGMENT - VEDANT SRIVASTAVA

Used Clean Architecture and Gin framework

API request Curl -
curl --location 'http://127.0.0.1:8080/event-process' \
--header 'Content-Type: application/json' \
--data '{
"ev": "top_cta_clicked",
"et": "clicked",
"id": "cl_app_id_001",
"uid": "cl_app_id_001-uid-001",
"mid": "cl_app_id_001-uid-001",
"t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"l": "en-US",
"sc": "1920 x 1080",
"atrk1": "button_text",
"atrv1": "Free trial",
"atrt1": "string",
"atrk2": "color_variation",
"atrv2": "ESK0023",
"atrt2": "string",
"atrk3": "page_path",
"atrv3": "/blog/category_one/blog_name.html",
"atrt3": "string",
"atrk4": "source",
"atrv4": "facebook",
"atrt4": "string",
"uatrk1": "user_score",
"uatrv1": "1034",
"uatrt1": "number",
"uatrk2": "gender",
"uatrv2": "m",
"uatrt2": "string",
"uatrk3": "tracking_code",
"uatrv3": "POSERK093",
"uatrt3": "string",
"uatrk4": "phone",
"uatrv4": "9034432423",
"uatrt4": "number",
"uatrk5": "coupon_clicked",
"uatrv5": "true",
"uatrt5": "boolean",
"uatrk6": "opt_out",
"uatrv6": "false",
"uatrt6": "boolean"
}'

Webhook Curl -
curl --location 'https://webhook.site/627bb56e-e865-48fb-a005-d4d63c193275' \
--header 'Content-Type: application/json' \
--data '{
"event": "top_cta_clicked",
"event_type": "clicked",
"app_id": "cl_app_id_001",
"user_id": "cl_app_id_001-uid-001",
"message_id": "cl_app_id_001-uid-001",
"page_title": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"page_url": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"browser_language": "en-US",
"screen_size": "1920 x 1080",
"attributes": {
"button_text": {
"value": "Free trial",
"type": "string"
},
"color_variation": {
"value": "ESK0023",
"type": "string"
},
"page_path": {
"value": "/blog/category_one/blog_name.html",
"type": "string"
},
"source": {
"value": "facebook",
"type": "string"
}
},
"traits": {
"coupon_clicked": {
"value": "true",
"type": "boolean"
},
"gender": {
"value": "m",
"type": "string"
},
"opt_out": {
"value": "false",
"type": "boolean"
},
"phone": {
"value": "9034432423",
"type": "number"
},
"tracking_code": {
"value": "POSERK093",
"type": "string"
},
"user_score": {
"value": "1034",
"type": "number"
}
}
}'
