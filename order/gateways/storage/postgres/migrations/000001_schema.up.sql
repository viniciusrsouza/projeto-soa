BEGIN;

CREATE TABLE IF NOT EXISTS orders (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    status text not null,
    ordered_by  int not null,
    property_id int not null,
    schedule_id int not null,
    property_owner_id int not null,
    created_at	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
