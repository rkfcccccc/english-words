import 'dart:convert';

class SourcedPicture {
  final String url;
  final String source;

  SourcedPicture({
    required this.url,
    required this.source,
  });

  factory SourcedPicture.fromMap(Map<String, dynamic> map) {
    return SourcedPicture(
      url: map['url'] ?? '',
      source: map['source'] ?? '',
    );
  }

  factory SourcedPicture.fromJson(String source) =>
      SourcedPicture.fromMap(json.decode(source));
}

class Definition {
  final String text;
  final String example;

  Definition({
    required this.text,
    required this.example,
  });

  factory Definition.fromMap(Map<String, dynamic> map) {
    return Definition(
      text: map['text'] ?? '',
      example: map['example'] ?? '',
    );
  }

  factory Definition.fromJson(String source) =>
      Definition.fromMap(json.decode(source));
}

class Meaning {
  final String partOfSpeech;
  List<Definition> definitions;
  List<String>? synonyms;
  List<String>? antonyms;

  Meaning({
    required this.partOfSpeech,
    required this.definitions,
    this.synonyms,
    this.antonyms,
  });

  factory Meaning.fromMap(Map<String, dynamic> map) {
    return Meaning(
      partOfSpeech: map['part_of_speech'] ?? '',
      definitions: List<Definition>.from(
          map['definitions']?.map((x) => Definition.fromMap(x))),
      synonyms:
          map['synonyms'] != null ? List<String>.from(map['synonyms']) : null,
      antonyms:
          map['antonyms'] != null ? List<String>.from(map['antonyms']) : null,
    );
  }

  factory Meaning.fromJson(String source) =>
      Meaning.fromMap(json.decode(source));
}

class WordEntry {
  final String id;
  final String word;
  final String phonetic;
  final List<Meaning> meanings;
  final List<SourcedPicture>? pictures;

  WordEntry({
    required this.id,
    required this.word,
    required this.phonetic,
    required this.meanings,
    this.pictures,
  });

  factory WordEntry.fromMap(Map<String, dynamic> map) {
    return WordEntry(
      id: map['id'] ?? '',
      word: map['word'] ?? '',
      phonetic: map['phonetic'] ?? '',
      meanings:
          List<Meaning>.from(map['meanings']?.map((x) => Meaning.fromMap(x))),
      pictures: map['pictures'] != null
          ? List<SourcedPicture>.from(
              map['pictures']?.map((x) => SourcedPicture.fromMap(x)))
          : null,
    );
  }

  factory WordEntry.fromJson(String source) =>
      WordEntry.fromMap(json.decode(source));
}
