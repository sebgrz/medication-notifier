import { useEffect, useState } from "react";
import styles from "./medicationsPanel.style.module.css";

export enum TimeOfDay {
	MORNING = "MOR",
	MIDDAY = "MID",
	EVENING = "EVE"
}

export enum Day {
	MONDAY = "MO",
	TUESDAY = "TU",
	WEDNESDAY = "WE",
	THURSDAY = "TH",
	FRIDAY = "FR",
	SATURDAY = "SA",
	SUNDAY = "SU"
}

export type Medication = {
	id: string,
	user_id: string,
	name: string,
	day: Day,
	time_of_day: TimeOfDay
}

const MedicationsPanel = (props: { data: Medication[] }) => {
	const [medications, setMedications] = useState<Medication[]>([]);

	useEffect(() => {
		setMedications(props.data);
	}, [props.data]);

	const getMedications = (day: Day, timeOfDay: TimeOfDay): Medication[] =>
		medications.filter(f => f.day === day && f.time_of_day === timeOfDay)

	return (
		<div className={styles.medicationPanel}>
			<table style={{ width: "100%" }}>
				<thead>
					<tr>
						<th>DAY</th>
						<th>Morning</th>
						<th>Midday</th>
						<th>Evening</th>
					</tr>
				</thead>
				<tbody>
					{Object.entries(Day).map(([k, v]) =>
						<tr key={k} style={{ border: "1px solid black" }}>
							<td>{k}</td>
							<td>{getMedications(v, TimeOfDay.MORNING).map(m => m.name)}</td>
							<td>{getMedications(v, TimeOfDay.MIDDAY).map(m => m.name)}</td>
							<td>{getMedications(v, TimeOfDay.EVENING).map(m => m.name)}</td>
						</tr>
					)}
				</tbody>
			</table>
		</div>
	);
}

export default MedicationsPanel;
