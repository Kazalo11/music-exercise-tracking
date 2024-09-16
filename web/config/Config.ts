import { promises as fs } from 'fs';
import * as yaml from 'js-yaml';

type Config = {
	server: {
		host: string;
	}
}

export async function loadConfig(filePath: string): Promise<Config> {
	try {
		const fileContents = await fs.readFile(filePath, 'utf8');
	    
		const data = yaml.load(fileContents) as Config;
	    
		return data;
	      } catch (err) {
		console.error(`Error reading YAML file: ${err}`);
		throw err;
	      }
}