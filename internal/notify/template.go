package notify

const Template = `{
    "user_id": "{{.receiver}}",
    "msg_type": "post",
    "content": {
        "post": {
            "zh_cn": {
                "title": "{{.title}}",
                "content": [
                    [
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "{{.text}}"
                        }
                    ],
                    [
                        {
                            "tag": "img",
                            "image_key": "{{.img}}",
                            "width": 300,
                            "height": 300
                        }
                    ],
                    [
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "PS: token有效期为30天, 小程序每10天更新"
                        }
                    ]
                ]
            }
        }
    }
}
`
