import 'package:english_words/utils/api/api.dart';

import 'models.dart';

Future<List<Movie>> search(String query) async {
  final data = await request(
    "get",
    "/movies",
    params: {"query": query},
    authRequired: true,
  );

  return List<Movie>.from(data.map((e) => Movie.fromMap(e)));
}

Future<void> favorite(int movieId) async {
  await request(
    "post",
    "/movies/$movieId/favorite",
    authRequired: true,
  );
}

Future<void> unfavorite(int movieId) async {
  await request(
    "delete",
    "/movies/$movieId/favorite",
    authRequired: true,
  );
}
