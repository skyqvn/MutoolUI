unit SettingUnit;

{$mode ObjFPC}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ComCtrls, StdCtrls,
  ExtCtrls;

type

  { TSettingForm }

  TSettingForm = class(TForm)
    CancelButton: TButton;
    OKButton: TButton;
    MainPageControl: TPageControl;
    Panel1: TPanel;
    MutoolPage: TTabSheet;
  private

  public

  end;

var
  SettingForm: TSettingForm;

implementation

{$R *.lfm}

end.

