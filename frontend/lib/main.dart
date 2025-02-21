import 'package:flutter/material.dart';
import 'package:moon_design/moon_design.dart';

import './pages/login.dart';

void main() {
  runApp(const App());
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      // NOTE: Need to find a better name for the app
      title: 'Local run',

      theme: ThemeData.light().copyWith(
        extensions: <ThemeExtension<dynamic>>[
          MoonTheme(tokens: MoonTokens.light),
        ],
      ),

      darkTheme: ThemeData.dark().copyWith(
        extensions: <ThemeExtension<dynamic>>[
          MoonTheme(tokens: MoonTokens.dark),
        ],
      ),

      initialRoute: '/login',
      routes: <String, WidgetBuilder>{
        '/login': (BuildContext context) => const Login(),
      },
    );
  }
}
