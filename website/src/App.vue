<script setup lang="ts">
import {RouterView} from 'vue-router'
</script>

<template>
  <div class="wrapper">
    <aside class="left-menu">
      <h1 class="title">Airport Dashboard</h1>

      <ul class="menu">
        <li class="menu-item">
          <RouterLink to="/" class="menu-link active">
            Dashboard
          </RouterLink>
        </li>
        <li class="menu-item">
          <RouterLink to="/others" class="menu-link">
            Autres
          </RouterLink>
        </li>
      </ul>
    </aside>

    <div class="right-part">
      <header>

      </header>
      <main>
        <RouterView/>
      </main>
    </div>
  </div>
</template>

<style scoped lang="scss">
.wrapper {
  height: 100vh;
  display: flex;

  .left-menu {
    width: 250px;
    background-color: #FFFFFF;
    border-right: 1px solid #F0F0F0;
    box-shadow: 0 0 10px rgba(0, 0, 0, .05);

    .title {
      height: 50px;
      font-size: 15px;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: #FFFFFF;
      border-bottom: 1px solid #F0F0F0;
    }

    .menu {
      list-style: none;
      padding: 0;
      margin: 0;

      .menu-item {
        width: 100%;

        .menu-link {
          text-decoration: none;
          color: #000;
          height: 50px;
          display: flex;
          align-items: center;
          justify-content: center;

          &.active {
            font-weight: bold;
          }
        }
      }
    }
  }

  .right-part {
    width: 100%;

    header {
      height: 50px;
      background-color: #FFFFFF;
      box-shadow: 0 0 10px rgba(0, 0, 0, .05);
      display: flex;
      align-items: center;
      padding: 10px;
    }
  }
}

</style>
