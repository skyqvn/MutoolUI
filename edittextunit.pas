unit EditTextUnit;

{$mode ObjFPC}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls;

type

  { TEditTextDialog }

  TEditTextDialog = class(TForm)
    CancelButton: TButton;
    OKButton: TButton;
    Panel1: TPanel;
    TextEdit: TEdit;
    Label1: TLabel;
  private

  public

  end;

var
  EditTextDialog: TEditTextDialog;

implementation

{$R *.lfm}

{ TEditTextDialog }

end.
