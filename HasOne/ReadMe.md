# Foreign key constraints

## Must insert table Account first, then insert table Student because column Student.AccountId links to primary key Account.AccountId

## Cannot insert rows into table Student if that row contains an AccountId that does not exist in table Account

## Only apply to Has-One relation, not available in Has-Many and Many-To-Many relations