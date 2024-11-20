<script>
  import {Zonetime} from '../wailsjs/go/main/App.js'
	function getZonetime() {
		Zonetime(zone).then((result) => zTime = result.split(' '))
	}
	let zTime = [];
	const zones = [
		'America/Los_Angeles', 
		'America/New_York',
		'Europe/London',
		'Europe/Paris',
		'Europe/Berlin',
		'Europe/Zurich',
		'Europe/Budapest',
		'Europe/Moscow',
		'Asia/Kabul',
		'Asia/Kolkata',
		'Asia/Tokyo',
		'Australia/Sydney',
		'Pacific/Auckland',
		'Pacific/Honolulu',
		'Pacific/Kiritimati',
	];
	$: thr = parseInt(zTime[4]);
	$: tmi = parseInt(zTime[5]);
	$: tse = parseInt(zTime[6]);

	$: angh = (thr*30 + tmi/2) - 90;
	$: angm = (tmi*6 + tse/10) - 90;

	$: handclr = thr < 6 ? '#666' : thr < 12 ? 'blue' : thr < 18 ? 'seagreen' : 'navy';

	let zone = zones[0]
	getZonetime()
</script>
<style>
	select {
		line-height: 2rem;
	}
</style>
<div class="my-4 flex gap-1 items center">
	<div class="text-3xl">Zone Times</div>
	<button class="p-1 rounded-md border border-gray-400" on:click={getZonetime}>Update</button>
</div>
<div class="grid grid-cols-3">
	<div class="col-span-1">
		<select size={zones.length} class="p-1 rounded-md text-gray-600 text-xl" bind:value={zone} on:change={getZonetime}>
			{#each zones as z}
			<option value={z}>{z.replace('/',':')}</option>					
			{/each}
		</select>
	</div>
	<div class="col-span-2">
		<div class="text-2xl">{zTime[0]} {zTime[1]} {zTime[2]} {zTime[4]}:{zTime[5]}:{zTime[6]} {zTime[7]}</div>
		<svg viewBox="-50 -50 100 100" width={400}>
			<circle r={49} fill="lightcyan" />
			<line x1={-6} x2={30} stroke={handclr} stroke-width={2} stroke-linecap="round" transform={`rotate(${angh})`} />
			<line x1={-6} x2={42} stroke={handclr} stroke-width={2} stroke-linecap="round" transform={`rotate(${angm})`} />
		</svg>
	</div>
</div>