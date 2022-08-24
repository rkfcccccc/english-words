import 'dart:convert';

import 'package:english_words/utils/api/dictionary/models.dart';

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
