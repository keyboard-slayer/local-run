import 'package:dio/dio.dart';

late Dio dio;

void initDio() {
  dio = Dio();
  dio.options.headers['X-LOCALRUN-CSRF-PROTECTION'] = '1';
  dio.options.headers['content-Type'] = 'application/json';
}
