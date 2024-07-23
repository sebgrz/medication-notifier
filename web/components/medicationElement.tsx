import { useApiManager } from "@/hooks/useApiManager";
import { Medication } from "./medicationsPanel"

const MedicationElement = (props: { element: Medication, removedMedicationAction: (m: Medication) => void }) => {
	const api = useApiManager();

	const removeMedicationOnClick = async () => {
		if (await api.appRemoveMedication(props.element.id)) {
			props.removedMedicationAction(props.element);
		}
	}
	return (
		<>
			<div>
				<button onClick={() => removeMedicationOnClick()}>❌</button>
				<span>{props.element.name}</span>
			</div>
		</>
	);
}

export default MedicationElement;
