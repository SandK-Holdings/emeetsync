import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

import '../router/app_router.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AutoTabsScaffold(
      routes: const [
        MySchejTabRoute(),
        EventsTabRoute(),
        FriendsTabRoute(),
        ProfileTabRoute(),
      ],
      bottomNavigationBuilder: _buildBottomNavigation,
    );
  }

  Widget _buildBottomNavigation(_, tabsRouter) {
    return BottomNavigationBar(
      items: <BottomNavigationBarItem>[
        BottomNavigationBarItem(
          icon: tabsRouter.activeIndex == 0
              ? const Icon(Icons.calendar_today)
              : const Icon(Icons.calendar_today_outlined),
          label: 'My schej',
        ),
        BottomNavigationBarItem(
          icon: tabsRouter.activeIndex == 1
              ? const Icon(MdiIcons.calendarMultipleCheck)
              : const Icon(MdiIcons.calendarBlankMultiple),
          label: 'My events',
        ),
        BottomNavigationBarItem(
          icon: tabsRouter.activeIndex == 2
              ? const Icon(Icons.people)
              : const Icon(Icons.people_outline),
          label: 'Friends',
        ),
        BottomNavigationBarItem(
          icon: tabsRouter.activeIndex == 3
              ? const Icon(Icons.person)
              : const Icon(Icons.person_outline),
          label: 'Profile',
        ),
      ],
      type: BottomNavigationBarType.fixed,
      currentIndex: tabsRouter.activeIndex,
      onTap: tabsRouter.setActiveIndex,
      backgroundColor: Colors.white,
      selectedItemColor: Colors.black,
      unselectedItemColor: Colors.grey,
    );
  }
}