<script>
    let users = [];
  
    async function fetchUsers() {
      const response = await fetch('http://localhost:8080/users');
      users = await response.json();
    }
  
    fetchUsers();
  </script>
  
  <main>
    <h1>Users</h1>
    {#each users as user}
      <div>
        <p>Username: {user.username}</p>
        <p>Email: {user.email}</p> 
      </div>
    {/each}
  </main>
  
