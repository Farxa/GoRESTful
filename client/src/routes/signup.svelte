<script>
    let username = '';
    let password = '';
    let avatar = ''; 
  
    async function handleSubmit(event) {
      event.preventDefault();
  
      try {
        const response = await fetch('http://localhost:8080/users', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ username, password, avatar })
        });
  
        if (!response.ok) {
          throw new Error(`Error: ${response.statusText}`);
        }
  
        const result = await response.json();
       console.log(result.username + ' has been created');
        // Redirect or clear form here if needed
      } catch (error) {
        console.error(error);
      }
    }
  </script>
  
  <div class="min-h-screen bg-gray-100 flex items-center justify-center">
    <form on:submit={handleSubmit} class="bg-white p-6 rounded shadow-md w-full max-w-sm">
      <div class="mb-4">
        <label for="username" class="block text-gray-700 text-sm mb-2">Username:</label>
        <input
          type="text"
          id="username"
          bind:value={username}
          required
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" />
      </div>
  
      <div class="mb-6">
        <label for="password" class="block text-gray-700 text-sm mb-2">Password:</label>
        <input
          type="password"
          id="password"
          bind:value={password}
          required
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" />
      </div>
  
      <div class="flex items-center justify-center">
        <button type="submit" class="bg-black hover:bg-gray-900 text-white py-2 px-4 rounded focus:outline-none focus:shadow-outline">
          Sign Up
        </button>
      </div>
    </form>
  </div>
  
  
   