<template>
  <div>
    <header>
      <!-- Navigation Bar -->
      <nav>
        <ul>
          <li>
            <NuxtLink to="/">
              Home
            </NuxtLink>
          </li>
          <li>
            <NuxtLink to="/login">
              Login
            </NuxtLink>
          </li>
          <li>
            <NuxtLink to="/signup">
              Signup
            </NuxtLink>
          </li>
        </ul>
      </nav>
    </header>

    <main>
      <!-- The page content will be injected here -->
      <NuxtPage />
    </main>

    <footer>
      <p>&copy; 2024 My Event Site</p>
    </footer>
  </div>
</template>

<script setup>
// You can add logic specific to the layout if needed
</script>
