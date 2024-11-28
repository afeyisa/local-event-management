<template>
  <ul class="space-y-4  pl-10">
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-search text-blue-500" />

      <NuxtLink
        to="/"
        class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
      >
        Browse Events
      </NuxtLink>
    </li>

    <li class="flex items-center space-x-2 ">
      <i class="fa fa-calendar text-blue-500" />

      <NuxtLink
        to="/events"
        class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
      >
        My Events
      </NuxtLink>
    </li>
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-bookmark text-blue-500" />
      <NuxtLink
        to="/events/bookmarks"
        class=" text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
      >
        Bookmarks
      </NuxtLink>
    </li>
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-user text-blue-500" />
      <NuxtLink
        to="/events/following"
        class=" text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
      >
        Following
      </NuxtLink>
    </li>
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-ticket text-blue-500" />
      <NuxtLink
        to="/events/mytickets"
        class=" text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
      >
        Tickets
      </NuxtLink>
    </li>

    <li class="flex items-center space-x-2 ">
      <i class="fa fa-building text-blue-500" />
      <NuxtLink
        to="/events/organizations"
        class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
      >
        Organizations
      </NuxtLink>
    </li>
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-plus text-blue-500" />

      <NuxtLink
        to="/events/create"
        class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500 "
      >
        Create Event
      </NuxtLink>
    </li>
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-plus text-blue-500" />

      <NuxtLink
        to="/events/organizations/create"
        class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500 "
      >
        Create Organizations
      </NuxtLink>
    </li>
    <li class="flex items-center space-x-2 ">
      <i class="fa fa-sign-out text-blue-500" />

      <DashboardLogoutButton />
    </li>
    <ThemeButton class="block sm:hidden" />
  </ul>
</template>
