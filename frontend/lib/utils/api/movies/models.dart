import 'dart:convert';

class Movie {
  final int id;
  final String imdbId;
  final String title;
  final int year;
  final String posterUrl;

  final double vocabularyPercent;

  Movie({
    required this.id,
    required this.imdbId,
    required this.title,
    required this.year,
    required this.posterUrl,
    required this.vocabularyPercent,
  });

  factory Movie.fromMap(Map<String, dynamic> map) {
    return Movie(
      id: map['id']?.toInt() ?? 0,
      imdbId: map['imdb_id'] ?? '',
      title: map['title'] ?? '',
      year: map['year']?.toInt() ?? 0,
      posterUrl: map['poster_url'] ?? '',
      vocabularyPercent: map['vocabulary_percent']?.toDouble() ?? 0.0,
    );
  }

  factory Movie.fromJson(String source) => Movie.fromMap(json.decode(source));
}
