import { Medication } from "./medicationsPanel"

const MedicationElement = (props: { element: Medication }) => {
	return (
		<>
			<div>
				<button>❌</button>
				<span>{props.element.name}</span>
			</div>
		</>
	);
}

export default MedicationElement;
