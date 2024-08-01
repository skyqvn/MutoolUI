unit SettingUnit;

{$mode ObjFPC}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ComCtrls, StdCtrls,
  ExtCtrls, ValEdit;

type

  { TSettingForm }

  TSettingForm = class(TForm)
    AddButton: TButton;
    SetAsActiveButton: TButton;
    EditButton: TButton;
    DeleteButton: TButton;
    CancelButton: TButton;
    Label1: TLabel;
    Label2: TLabel;
    AllListBox: TListBox;
    OKButton: TButton;
    MainPageControl: TPageControl;
    Panel1: TPanel;
    MutoolPage: TTabSheet;
    Panel2: TPanel;
    ActiveEditPanel: TPanel;
    Panel3: TPanel;
    Panel4: TPanel;
  private

  public

  end;

var
  SettingForm: TSettingForm;

implementation

{$R *.lfm}

{ TSettingForm }


end.

