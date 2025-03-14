# Financial Assistance Schema Management System

**Creator:**  
Hoo Di Heng  

**Back-End:**  
GoLang, Gin, GORM, validator, uuid

### API Endpoints 
**Applications**  
GET  /api/applications -> Get all applications  
GET  /api/application/{ID} -> Get application by id  
POST  /api/applications -> Create a new application  
DELETE  /api/applications/{ID} -> Delete an application by id  
DELETE  /api/applications -> Delete all applications

**Applicant**  
GET  /api/applicants -> Get all applicant  
GET  /api/applicant/{ID} -> Get applicant by id  
POST  /api/applicants -> Create a new applicant  
DELETE  /api/applicants/{ID} -> Delete an applicant by id  
DELETE  /api/applicants -> Delete all applicants

**Schemes**  
GET  /api/schemes -> Get all schemes  
GET  /api/schemes/{ID} -> Get scheme by id  
POST  /api/schemes -> Create a new schemes  
DELETE  /api/schemes/{ID} -> Delete a scheme by id  
DELETE  /api/schemes -> Delete all schemes