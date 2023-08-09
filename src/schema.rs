diesel::table! {
    users (id) {
        id -> Int4,
        name -> Varchar,
    }
}

diesel::table! {
    members (id) {
        id -> Int4,
        student_id -> VarChar,
        student_email -> VarChar,
        first_name -> VarChar,
        last_name -> VarChar,
        active_member -> Bool
    }
}

diesel::table! {
    admins (token) {
        token -> VarChar,
    }
}