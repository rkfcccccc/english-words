import 'dart:convert';

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
  WordEntry({
    required this.id,
    required this.word,
    required this.phonetic,
    required this.meanings,
  });

  factory WordEntry.fromMap(Map<String, dynamic> map) {
    return WordEntry(
      id: map['id'] ?? '',
      word: map['word'] ?? '',
      phonetic: map['phonetic'] ?? '',
      meanings:
          List<Meaning>.from(map['meanings']?.map((x) => Meaning.fromMap(x))),
    );
  }

  factory WordEntry.fromJson(String source) =>
      WordEntry.fromMap(json.decode(source));
}

class Challenge {
  final WordEntry entry;
  final int learningStep;

  Challenge({
    required this.entry,
    required this.learningStep,
  });

  factory Challenge.fromMap(Map<String, dynamic> map) {
    return Challenge(
      entry: WordEntry.fromMap(map['entry']),
      learningStep: map['learning_step']?.toInt() ?? 0,
    );
  }

  factory Challenge.fromJson(String source) =>
      Challenge.fromMap(json.decode(source));

  Challenge copyWith({
    WordEntry? entry,
    int? learningStep,
  }) {
    return Challenge(
      entry: entry ?? this.entry,
      learningStep: learningStep ?? this.learningStep,
    );
  }
}
