import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:english_words/utils/api/user/auth.dart';
import 'package:english_words/utils/errors.dart';

const authority = "localhost:8080";
const basePath = "/api";

final client = HttpClient();

Future<String> readResponse(HttpClientResponse response) {
  final completer = Completer<String>();
  final contents = StringBuffer();

  response.transform(utf8.decoder).listen((data) {
    contents.write(data);
  }, onDone: () {
    completer.complete(contents.toString());
  });

  return completer.future;
}

dynamic request(
  String httpMethod,
  String path, {
  Map<String, dynamic>? params,
  Map<String, dynamic>? body,
  authRequired = false,
}) async {
  final uri = Uri.http(authority, basePath + path, params);

  while (true) {
    try {
      final request = await client.openUrl(httpMethod, uri);

      while (refreshLock) {
        await Future.delayed(const Duration(milliseconds: 100));
        continue;
      }

      request.headers.add("accept", ContentType.json);

      if (authRequired) {
        final accessToken = await getAccessToken();
        request.headers.add("Authorization", "Bearer ${accessToken!}");
      }

      if (body != null) {
        request.headers.contentType = ContentType.json;
        request.write(jsonEncode(body));
      }

      final response = await request.close();
      if (response.statusCode >= 500) {
        throw const AppError(
          key: FailureError.serverInternalError,
          message: "Server internal error",
        );
      }

      dynamic data;
      final contentType = response.headers.contentType;
      if (contentType != null && contentType.mimeType == "application/json") {
        data = jsonDecode(await readResponse(response));
      } else if (contentType != null) {
        throw Exception(
          "got unexpected content type: ${response.headers.contentType}",
        );
      }

      if (response.statusCode == HttpStatus.unauthorized &&
          data?["error_name"] == "TOKEN_EXPIRED") {
        await refresh();
        await Future.delayed(const Duration(milliseconds: 250));
        continue;
      }

      if (response.statusCode >= 400) {
        throwSomeAppError(data?["error_name"]);
      }

      return data;
    } on SocketException {
      throw const AppError(
        key: FailureError.noInternet,
        message: "No Internet Connection",
      );
    }
  }
}
