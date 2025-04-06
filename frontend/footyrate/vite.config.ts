import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		allowedHosts: [
			'localhost',
			'127.0.0.1',
			'5a2b-111-93-74-158.ngrok-free.app'
		]
	}

});
