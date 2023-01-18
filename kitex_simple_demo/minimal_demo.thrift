namespace go minimalDemo

struct AddRequest {
  1: i64 a
  2: i64 b
}

struct AddResponse {
  1: i64 res
}

service AddService {
    AddResponse Add(1: AddRequest req)
}