<script>
  import logo from './assets/images/logo-universal.png'
  import {Greet} from '../wailsjs/go/main/App.js'
  import {Addnums} from '../wailsjs/go/main/App.js'
  import Clock from './Clock.svelte';

  let resultText = "Please enter your name below 👇"
  let name
	let nums = '';
	let sumOfnums = 0;

  function greet() {
		Greet(name).then(result => {
			resultText = result
		})
  }

	function calcSum() {
		Addnums(nums.split(',').map(n => parseInt(n))).then(result => sumOfnums = result)
	}
</script>

<main>
	<nav class="flex gap-1 items-center">
		<img alt="Wails logo" class="w-24 h-24" src="{logo}">
		<div class="text-4xl ml-4">Wails Demo App</div>
	</nav>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="p-1 bg-gray-300 text-black" id="name" type="text"/>
    <button class="rounded-md p-1 text-black bg-gray-300" on:click={greet}>Greet</button>
  </div>
	<fieldset class="my-4 border-2 border-gray-500 p-1">
		<legend class="text-2xl">Sum of numbers:</legend>
		<div>
			Enter numbers separated by comma: 
			<input class="p-1 bg-gray-300 text-black" type="search" bind:value={nums} on:input={calcSum} />
			<small>(click &times; to clear)</small>
		</div>

		<div>
			Sum of {nums.replaceAll(',','+')} = <span class="text-2xl">{sumOfnums}</span>
		</div>
	</fieldset>
	<div>
		<Clock />
	</div>
</main>

<style>

  #logo {
    display: block;
    width: 200px;
    height: 200px;
    margin: auto;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
