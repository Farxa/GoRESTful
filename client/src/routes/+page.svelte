<script>
  import Signup from "./signup.svelte";
  import "../app.css";

    let users = [];
  
    async function fetchUsers() {
      const response = await fetch('http://localhost:8080/users');
      users = await response.json();
    }
  
    fetchUsers();
  </script>
  
  <main>
    <Signup />
  </main>
  
