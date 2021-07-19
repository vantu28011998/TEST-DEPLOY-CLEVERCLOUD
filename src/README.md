# DlgO
This is a project from DacTech internship's GOTEAM that simulate DglO Project
## Installation
## Usage
## Design
### Folder struture
    <GOTEAM_DglO_BE>/
        | -- Clients/
           [ conversate with a external system which same as connect to database, line, facebook, google,... ]
        | -- cmd/
            | -- messenger/
               [ entry point of a program  ]
        | -- handlers/
           [ contain Restful APIs  ] 
        | -- middlewares/
           [ Authenticate request same as filter, when a request goes through which must be allowed  ]
        | -- models/
           [ contain data ]
        | -- services/
           [ bussiness logic ]
### Database design
### Restful API design
   #### MESSAGES APIs

   #### POST /api/v1/messages/text
   {
      "title":"dglo",
      "content":"dlgo",
      "status": 0
   }
   Note: status = 0 - Draft
         status = 1 - OK
   #### GET /api/v1/messages?title=dglo&type=[0]&page=1&limit=10
   #### SERVER RETURNS:
   {
      "pagination":{
         "page":1,
         "limit":10,
         "sumOfPage":25
      },
      messages:[
         {
            "id":1,
            "title":"dglo 1",
            "content":"dglo 1",
            "type":0,
            "status": 0,
            "createAt":"07/15/2021 13:10"
         },
         {  
            "id":2,
            "title":"dglo 2",
            "content":"dglo 2",
            "type":0,
            "status": 1
            "createAt":"07/15/2021 13:15"
         },
         ...
         {
            "id":10,
            "title":"dglo 10",
            "content":"dglo 10",
            "type":0,
            "status": 10,
             "createAt":"07/15/2021 13:20"
         },
      ]
   }
   Note: Type = 0 - Text
         Type = 1 - Imagemap

   #### MESSAGE SCHEDULES APIs

   #### POST /api/v1/schedules
   {
      "scheduleTitle":"DlgO",
      "sendingDatetime":"07/15/2021 15:05",
      "targetSegment":[
         15,
         7,
         2021
      ],
      "messages":[
        1,
        3,
        5,
        7,
        9
      ],
      "status":1,
   }
   #### GET /api/v1/schedules?title=dglo&type=[0]&status=[1]&fromAt=07/15/2021&to=08/02/2021&page=1&limit=10
   #### SERVER RETURNS:
   {
      "pagination":{
         "page":1,
         "limit":10,
         "sumOfPage":25
      },
      schedules:[
        {
         "id":1,
         "scheduleTitle":"DlgO",
         "sendingDatetime":"07/15/2021 15:05",
         "targetSegment":[
            15,
            7,
            2021
         ],
         "status":1,
         "createBy":"username"       
        },
         {
         "id":2,
         "scheduleTitle":"DlgO",
         "sendingDatetime":"07/15/2021 15:10",
         "targetSegment":[
            15,
            7,
            2021
         ],
         "status":1,       
        }              
      ]
   }


## Contributing
## License
DacTech internship's GOTEAM
