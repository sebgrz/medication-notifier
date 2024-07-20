'use client';

import { useFormik } from "formik";
import EnumSelector from "./enumSelector";
import { Day, Medication, TimeOfDay } from "./medicationsPanel";
import { useApiManager } from "@/hooks/useApiManager";

const AddMedicationPanel = (props: { addedMedicationsAction: (m: Medication[]) => void }) => {
	const api = useApiManager();
	const formik = useFormik({
		initialValues: {
			morning: '',
			midday: '',
			evening: ''
		},
		onSubmit: async (values) => {
			console.info(values);
			await api.appAddMedication(values.morning, Day.SATURDAY, TimeOfDay.MORNING);
			formik.resetForm();
		},
	});

	const addMedicationSelectDayOnChange = (day: typeof Day) => {
		console.info(day);
	}

	return (
		<tr>
			<th>
				<button type="submit" onClick={() => formik.submitForm()}>ADD</button>&nbsp;
				<EnumSelector enumType={Day} onChange={addMedicationSelectDayOnChange} />
			</th>
			<th><input type="text" id="morning" name="morning" onChange={formik.handleChange} value={formik.values.morning} /></th>
			<th><input type="text" id="midday" name="midday" onChange={formik.handleChange} value={formik.values.midday} /></th>
			<th><input type="text" id="evening" name="evening" onChange={formik.handleChange} value={formik.values.evening} /></th>
		</tr>
	);
}

export default AddMedicationPanel;
