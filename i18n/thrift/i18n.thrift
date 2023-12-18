
struct ListLanguagesRequest {
  1: list<string> languages,
}

struct ListLanguagesResponse {
  1: map<string, LanguageEntry> entries,
  2: i64 timestamp,
}

struct LanguageEntry {
    1: string path,
    2: string language,
    3: bool valid,
    20: binary payload,
}

// The I18N Thrift service definition.
service I18N {
  // ListLanguages
  ListLanguagesResponse ListLanguages (1: ListLanguagesRequest req);
}