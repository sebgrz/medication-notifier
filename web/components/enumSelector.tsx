import { ChangeEvent } from "react";

const EnumSelector = <TEnum,>(props: { enumType: any, onChange: (v: TEnum) => void }) => {
	const onChange = (event: ChangeEvent<HTMLSelectElement>) => {
		props.onChange((event.target.value as any) as TEnum);
	}

	return (
		<select onChange={onChange}>
			{
				Object.entries(props.enumType as object).map(([k, v]) => <option key={k} value={v}>{k}</option>)
			}
		</select>
	);
}

export default EnumSelector;
