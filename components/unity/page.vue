<template>
  <NavBar :page-title="pageTitle"></NavBar>
  <slot></slot>
  <QuickNav></QuickNav>
</template>

<script setup>
defineProps(["pageTitle"])
</script>