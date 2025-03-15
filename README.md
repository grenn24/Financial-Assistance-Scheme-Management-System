# Financial Assistance Schema Management System

Backend API that can perform CRUD operations on financial assistance schemes, applicants and applications
Returns an list of schemes an applicant is eligible for applying
Stores the state of the scheme application (pending, approved, rejected)

**Creator:**  
Hoo Di Heng

**Back-End:**  
GoLang, Gin, GORM, validator, uuid

### API Endpoints

**Applications**  
GET /api/applications -> Get all applications  
GET /api/applications/{ID} -> Get application by id  
POST /api/applications -> Create a new application  
PUT /api/applications/{ID} -> Update application details  
DELETE /api/applications/{ID} -> Delete an application by id  
DELETE /api/applications -> Delete all applications

**Applicant**  
GET /api/applicants -> Get all applicant  
GET /api/applicants/{ID} -> Get applicant by id  
POST /api/applicants -> Create a new applicant  
PUT /api/applicants/{ID} -> Update applicant details  
DELETE /api/applicants/{ID} -> Delete an applicant by id  
DELETE /api/applicants -> Delete all applicants

**Schemes**  
GET /api/schemes -> Get all schemes  
GET /api/schemes/{ID} -> Get scheme by id  
GET /api/schemes/search?query={name} -> Search for schemes by name  
GET /api/schemes/eligible?applicant={ID} -> Get eligible schemes for applicant  
POST /api/schemes -> Create a new schemes  
PUT /api/schemes/{ID} -> Update scheme details  
DELETE /api/schemes/{ID} -> Delete a scheme by id  
DELETE /api/schemes -> Delete all schemes

**Database Schema**
###Applicant / Household Member
Name
EmploymentStatus
MaritalStatus
Sex
DOB
Relation (household member only)
SchoolLevel (household member only)
HouseholdOwnerID (household member only)

###Application
SchemeID (fkey)
ApplicantID (fkey)
Status

###Scheme
Name
Description
Benefits (fkey)
Criteria (fkey)

**Response JSON Structure (GET)**
```json
// Applicant
{
	"id": "01959522-9d21-70f2-9b26-de4efafdd6d0",
	"created_at": "2025-03-14T22:50:28.513062+08:00",
	"updated_at": "2025-03-14T22:50:28.513062+08:00",
	"name": "Nathaniel",
	"employment_status": false,
	"marital_status": "married",
	"sex": "male",
	"date_of_birth": "2006-01-02T23:04:05+08:00",
	"household": [
		{
			"id": "01959522-9d23-7120-baf0-72f4a68734c1",
			"created_at": "2025-03-14T22:50:28.514315+08:00",
			"updated_at": "2025-03-14T22:50:28.514315+08:00",
			"household_owner_id": "01959522-9d21-70f2-9b26-de4efafdd6d0",
			"name": "Isabelle",
			"employment_status": false,
			"marital_status": "single",
			"sex": "female",
			"date_of_birth": "2006-01-02T23:04:05+08:00",
			"relation": "daughter",
            "school_level":"primary"
		}
	]
}

// Scheme
{
    "id": "01959558-a7bb-7574-92d5-7ae6f528671b",
    "created_at": "2025-03-14T23:49:30.171357+08:00",
    "updated_at": "2025-03-14T23:49:30.171357+08:00",
    "name": "Retrenchment Assistance Scheme",
    "description": "Financial assistance for retrenched workers",
    "benefits": [
        {
            "id": "01959558-a7bc-7626-9883-e001b0fae142",
            "created_at": "2025-03-14T23:49:30.172403+08:00",
            "updated_at": "2025-03-14T23:49:30.172403+08:00",
            "SchemeID": "01959558-a7bb-7574-92d5-7ae6f528671b",
            "name": "Additional SkillsFuture credits",
            "amount": 500.25
        }
    ],
    "criteria": {
        "id": "01959558-a7bc-7e0e-b7ad-5e3e8057da81",
        "created_at": "2025-03-14T23:49:30.17188+08:00",
        "updated_at": "2025-03-14T23:49:30.17188+08:00",
        "SchemeID": "01959558-a7bb-7574-92d5-7ae6f528671b",
        "employment_status": true,
        "marital_status": "single",
        "has_children": {
            "school_level": "all"
        }
    }
}

// Application
{
    "id": "01959568-ba49-79d9-9988-e4c665c8686e",
    "applicant_id": "01959522-9d21-70f2-9b26-de4efafdd6d0",
    "scheme_id": "01959558-a7bb-7574-92d5-7ae6f528671b",
    "status": "pending",
    "created_at": "2025-03-15T00:07:03.497645+08:00",
    "updated_at": "2025-03-15T00:07:03.497645+08:00",
    "applicant": {
        "id": "01959522-9d21-70f2-9b26-de4efafdd6d0",
        "created_at": "2025-03-14T22:50:28.513062+08:00",
        "updated_at": "2025-03-14T22:50:28.513062+08:00",
        "name": "Nathaniel",
        "employment_status": false,
        "marital_status": "married",
        "sex": "male",
        "date_of_birth": "2006-01-02T23:04:05+08:00",
        "household": [
            {
                "id": "01959522-9d22-74d0-aa3b-ac08ea5d3497",
                "created_at": "2025-03-14T22:50:28.514315+08:00",
                "updated_at": "2025-03-14T22:50:28.514315+08:00",
                "household_owner_id": "01959522-9d21-70f2-9b26-de4efafdd6d0",
                "name": "Isabelle",
                "employment_status": false,
                "marital_status": "single",
                "sex": "male",
                "date_of_birth": "2006-01-02T23:04:05+08:00",
                "relation": "daughter",
                "school_level":"primary"
            },
            {
                "id": "01959522-9d23-7120-baf0-72f4a68734c1",
                "created_at": "2025-03-14T22:50:28.514315+08:00",
                "updated_at": "2025-03-14T22:50:28.514315+08:00",
                "household_owner_id": "01959522-9d21-70f2-9b26-de4efafdd6d0",
                "name": "Isabelle",
                "employment_status": false,
                "marital_status": "single",
                "sex": "male",
                "date_of_birth": "2006-01-02T23:04:05+08:00",
                "relation": "daughter"
            }
        ]
    },
    "scheme": {
        "id": "01959558-a7bb-7574-92d5-7ae6f528671b",
        "created_at": "2025-03-14T23:49:30.171357+08:00",
        "updated_at": "2025-03-14T23:49:30.171357+08:00",
        "name": "Retrenchment Assistance Scheme",
        "description": "Financial assistance for retrenched workers",
        "benefits": [
            {
                "name": "Additional SkillsFuture credits",
                "amount": 500.25
            },
            {
                "name": "Additional SkillsFuture credits",
                "amount": 500.25
            }
        ],
        "criteria": {
            "employment_status": true,
            "marital_status": "single",
            "has_children": {
                "school_level": "all"
            }
        }
    }
}
```

**Request JSON Structure (POST)**
```json
// Applicant
{
    "name": "Nathaniel",
    "employment_status": false,
    "marital_status":"single",
    "sex":"male",
    "date_of_birth":"2006-01-02T15:04:05Z",
    "household":[{
        "name": "Isabelle",
        "employment_status": false,
        "marital_status":"single",
        "sex":"female",
        "date_of_birth":"2006-01-02T15:04:05Z",
        "relation":"daughter",
        "school_level":"primary"
    },
    {
        "name": "Isabelle",
        "employment_status": false,
        "marital_status":"single",
        "sex":"female",
        "date_of_birth":"2006-01-02T15:04:05Z",
        "relation":"daughter",
    }]
}

// Scheme
{
    "name": "Retrenchment Assistance Scheme",
    "description":"Financial assistance for retrenched workers",
    "criteria":{
        "employment_status":true,
        "marital_status":"single",
        "has_children":{
            "school_level":"all"
        }
    },
    "benefits":[
        {
            "name":"Additional SkillsFuture credits",
            "amount":500.25
        }
    ]
}

// Application
{
    "applicant_id":"01959522-9d21-70f2-9b26-de4efafdd6d0",
    "scheme_id":"01959558-a7bb-7574-92d5-7ae6f528671b",
    "status":"pending"
}
```