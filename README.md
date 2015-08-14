# sendgrid-activities

## Usage

```
Usage of sendgrid-activities:
  -password="REQUIRED": Password to connect to the SendGrid API
  -type="all": Types of activities to retrieve (all, bounces, blocks, invalidEmails)
  -username="REQUIRED": Username to connect to the SendGrid API
```

**Example**
```
sendgrid-activities -username=sendgridusername -password=P@ssw0rd1 -type=all
```

**Output to CSV**
```
sendgrid-activities -username=sendgridusername -password=P@ssw0rd1 > activities.csv
```
