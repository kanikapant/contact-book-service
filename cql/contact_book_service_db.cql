// To execute CQL statements from a file use SOURCE 'file_name.cql'
// To execute a CQL file without starting a shell session use cqlsh [-u user -p pwd] -f 'file_name.cql'

CREATE KEYSPACE IF NOT EXISTS contact_book_service_db
WITH REPLICATION = {
	'class': 'SimpleStrategy',
	'replication_factor': '1'}
AND DURABLE_WRITES = true;