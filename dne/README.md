DNE - by WaGal

This app enriches the brazilian National Address Directory (DNE), a list of brazilian postal codes with their base addresses. 
DNE itens are on address table in the database.

Brazilian Post Office does not deliver mail in some postal codes. For these postal codes the mail receivers must go to the nearest 
Post Office Station to get their mail. Brazilian Post Office specifies ranges of restricted postal codes. Each range is a row in 
restriction table on the database. This app takes each range and to each postal code in the range updates the corresponding postal 
code on address table with the restriction.

Besides updating all restrictions on address table, this app takes each postal code on address table and tries to get the postal 
code base address geographic coordinates (latitude and longitude), using Bing Geocode Rest Api, and updates the postal code on 
address table with that info.

Finally, the address table is made available for searching by registered users. The user, after signing in, just have to fill a 
form, typing the postal code number and choosing the response format between html, xml or json.

Just edit the files wagnergaldino/training/dne/database/database.go and wagnergaldino/training/dne/orm/orm.go to update the database connection info. There's a mysql database structure dump available at wagnergaldino/training/dne/database folder. 

The app runs on port 8081 and user/pass = wagal/wagal

Enjoy !!!   ;-)
