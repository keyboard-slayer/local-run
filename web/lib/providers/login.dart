import 'package:flutter_dotenv/flutter_dotenv.dart';
import '../utils/dio.dart';

void signIn(String username, String password) async {
  String? backend = dotenv.env['BACKEND_URL'];

  if (backend == null) {
    throw Exception('The backend is not set');
  }

  var response = await dio.post(
    '$backend/auth',
    data: {'username': username, 'password': password},
  );

  if (response.statusCode != 200) {
    throw Exception("Unexpected backend error: ${response.statusMessage}");
  }
}

