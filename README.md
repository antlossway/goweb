# goweb
Build web service with Go


#emailtool

A web portal to send group emails to customers. It has predefined 2 tabs, one to send promotion email, one to send
notifications email. You can define separate sets of contact lists for the 2 tabs. You can as well add more tabs for more usage case.

## configration
under cfg directory, there are several files:
- email.cfg: set up general email field, including "From", "Cc", and Signature
- smtp.cfg: parameters to connect to SMTP server
- promotion_contact_json.cfg: contact list to send promotion in json format
- notification_contact_json.cfg: contact list to send notificactions (e.g server maintenance, outage..etc)



