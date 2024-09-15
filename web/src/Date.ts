import { format, parse, parseISO } from "date-fns";

export function formatString(date: string) {
	const parsedDate = parse(date, "dd MMM yyyy, HH:mm", new Date());
	return format(parsedDate, "MMMM dd, yyyy h:mm a");
}

export function formatISOString(isoDate: string) {
	const date = parseISO(isoDate);
	return format(date, "MMMM dd, yyyy h:mm a");
      
}